import axios from "axios";

const ACCESS_TOKEN_KEY = "accessToken";
const REFRESH_TOKEN_KEY = "refreshToken";

type RefreshResponse = {
  accessToken: string;
};

type RetryableRequest = {
  _retry?: boolean;
};

let refreshAccessTokenPromise: Promise<string> | null = null;

function getAccessToken(): string | null {
  return localStorage.getItem(ACCESS_TOKEN_KEY);
}

function getRefreshToken(): string | null {
  return localStorage.getItem(REFRESH_TOKEN_KEY);
}

export function saveAuthTokens(accessToken: string, refreshToken: string) {
  localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
  localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
}

export function clearAuthTokens() {
  localStorage.removeItem(ACCESS_TOKEN_KEY);
  localStorage.removeItem(REFRESH_TOKEN_KEY);
}

function redirectToLogin() {
  if (window.location.pathname !== "/login") {
    window.location.replace("/login");
  }
}

async function refreshAccessToken(): Promise<string> {
  if (refreshAccessTokenPromise !== null) {
    return refreshAccessTokenPromise;
  }

  const refreshToken = getRefreshToken();
  if (refreshToken === null) {
    clearAuthTokens();
    redirectToLogin();
    throw new Error("missing refresh token");
  }

  refreshAccessTokenPromise = axios
    .post<RefreshResponse>("/api/v1/auth/refresh", {
      refreshToken,
    })
    .then((response) => {
      const nextAccessToken = response.data.accessToken;
      localStorage.setItem(ACCESS_TOKEN_KEY, nextAccessToken);
      return nextAccessToken;
    })
    .catch((error) => {
      if (axios.isAxiosError(error) && error.response?.status === 401) {
        clearAuthTokens();
        redirectToLogin();
      }

      throw error;
    })
    .finally(() => {
      refreshAccessTokenPromise = null;
    });

  return refreshAccessTokenPromise;
}

export const httpClient = axios.create();

httpClient.interceptors.request.use((config) => {
  const accessToken = getAccessToken();
  if (accessToken !== null) {
    config.headers.Authorization = `Bearer ${accessToken}`;
  }

  return config;
});

httpClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (!axios.isAxiosError(error) || error.response?.status !== 401 || error.config === undefined) {
      return Promise.reject(error);
    }

    const requestUrl = error.config.url ?? "";
    const isLoginRequest = requestUrl.includes("/api/v1/auth/login");
    const isSignupRequest = requestUrl.includes("/api/v1/auth/signup");
    const isRefreshRequest = requestUrl.includes("/api/v1/auth/refresh");

    if (isRefreshRequest) {
      clearAuthTokens();
      redirectToLogin();
      return Promise.reject(error);
    }

    if (isLoginRequest || isSignupRequest) {
      return Promise.reject(error);
    }

    const retryableRequest = error.config as typeof error.config & RetryableRequest;
    if (retryableRequest._retry) {
      return Promise.reject(error);
    }

    retryableRequest._retry = true;

    try {
      const nextAccessToken = await refreshAccessToken();
      retryableRequest.headers.Authorization = `Bearer ${nextAccessToken}`;
      return httpClient(retryableRequest);
    } catch (refreshError) {
      return Promise.reject(refreshError);
    }
  },
);
