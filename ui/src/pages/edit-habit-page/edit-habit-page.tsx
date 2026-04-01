import type { FormEvent } from "react";
import axios from "axios";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { IconSelector } from "../../components/icon-selector/icon-selector";
import type { HabitDto, UpdateHabitDto, ValidationErrorsDto } from "../../dto/habit";
import { httpClient } from "../../infrastructure/http-client/http-client";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./edit-habit-page.sass";

export function EditHabitPage() {
  const navigate = useNavigate();
  const { uuid = "" } = useParams();

  const [icon, setIcon] = useState(100);
  const [name, setName] = useState("");
  const [fieldErrors, setFieldErrors] = useState<ValidationErrorsDto>({});
  const [requestError, setRequestError] = useState("");
  const [loadedHabit, setLoadedHabit] = useState<HabitDto | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  useEffect(() => {
    let isMounted = true;

    async function loadHabit() {
      setIsLoading(true);
      setRequestError("");

      try {
        const response = await httpClient.get<HabitDto>(`/api/v1/journal/habits/${uuid}`);

        if (!isMounted) {
          return;
        }

        setLoadedHabit(response.data);
        setIcon(response.data.icon);
        setName(response.data.name);
      } catch (error) {
        if (!isMounted) {
          return;
        }

        if (axios.isAxiosError(error) && typeof error.response?.data === "object" && error.response?.data !== null) {
          const responseData = error.response.data as Record<string, unknown>;

          if ("message" in responseData && typeof responseData.message === "string") {
            setRequestError(responseData.message);
          } else {
            setRequestError("Failed to load habit.");
          }
        } else {
          setRequestError("Failed to load habit.");
        }
      } finally {
        if (isMounted) {
          setIsLoading(false);
        }
      }
    }

    if (uuid !== "") {
      void loadHabit();
    } else {
      setRequestError("Habit uuid is missing.");
      setIsLoading(false);
    }

    return () => {
      isMounted = false;
    };
  }, [uuid]);

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setFieldErrors({});
    setRequestError("");

    try {
      const payload: UpdateHabitDto = {
        icon,
        name,
      };

      const response = await httpClient.put<HabitDto>(`/api/v1/journal/habits/${uuid}`, payload);
      setLoadedHabit(response.data);
      setIcon(response.data.icon);
      setName(response.data.name);
    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.data !== undefined) {
        if (typeof error.response.data === "object" && error.response.data !== null) {
          const responseData = error.response.data as Record<string, unknown>;

          if ("message" in responseData && typeof responseData.message === "string") {
            setRequestError(responseData.message);
          } else {
            setFieldErrors(error.response.data as ValidationErrorsDto);
          }
        } else {
          setRequestError("Request failed. Try again.");
        }
      } else {
        setRequestError("Request failed. Try again.");
      }
    } finally {
      setIsSubmitting(false);
    }
  }

  async function onDelete() {
    if (!window.confirm("Delete this habit?")) {
      return;
    }

    setIsDeleting(true);
    setRequestError("");

    try {
      await httpClient.delete(`/api/v1/journal/habits/${uuid}`);
      navigate("/habits");
    } catch (error) {
      if (axios.isAxiosError(error) && typeof error.response?.data === "object" && error.response?.data !== null) {
        const responseData = error.response.data as Record<string, unknown>;

        if ("message" in responseData && typeof responseData.message === "string") {
          setRequestError(responseData.message);
        } else {
          setRequestError("Failed to delete habit.");
        }
      } else {
        setRequestError("Failed to delete habit.");
      }
    } finally {
      setIsDeleting(false);
    }
  }

  return (
    <AdaptiveLayout
      eyebrow="Acta / Edit Habit"
      title="Update the habit details or remove it from the active list."
      description="This page loads the current habit from the journal API and writes changes back through the update route."
    >
      <section className="edit-habit-page__grid">
        <article className="edit-habit-page__card edit-habit-page__card--form">
          <span className="edit-habit-page__eyebrow">Edit habit</span>
          <h2 className="edit-habit-page__title">Habit settings</h2>

          {isLoading ? (
            <p className="edit-habit-page__text">Loading habit...</p>
          ) : (
            <form className="edit-habit-page__form" onSubmit={onSubmit}>
              <label className="edit-habit-page__field">
                <span className="edit-habit-page__label">Name</span>
                <input
                  className="edit-habit-page__input"
                  name="name"
                  type="text"
                  value={name}
                  onChange={(event) => setName(event.target.value)}
                />
              </label>

              <IconSelector label="Icon" value={icon} onChange={setIcon} />

              <div className="edit-habit-page__actions">
                <button className="edit-habit-page__submit" disabled={isSubmitting || isDeleting} type="submit">
                  {isSubmitting ? "Saving..." : "Save habit"}
                </button>
                <button
                  className="edit-habit-page__delete"
                  disabled={isSubmitting || isDeleting}
                  type="button"
                  onClick={() => void onDelete()}
                >
                  {isDeleting ? "Deleting..." : "Delete habit"}
                </button>
              </div>
            </form>
          )}

          {Object.keys(fieldErrors).length > 0 && (
            <ul className="edit-habit-page__errors">
              {Object.entries(fieldErrors).map(([field, message]) => (
                <li className="edit-habit-page__error" key={`${field}-${message}`}>
                  <span className="edit-habit-page__error-field">{field}</span>
                  <strong className="edit-habit-page__error-message">{message}</strong>
                </li>
              ))}
            </ul>
          )}

          {requestError !== "" && <div className="edit-habit-page__request-error">{requestError}</div>}
        </article>

        <article className="edit-habit-page__card edit-habit-page__card--summary">
          <span className="edit-habit-page__eyebrow">Current state</span>
          <h2 className="edit-habit-page__title">Loaded habit</h2>

          {loadedHabit === null ? (
            <p className="edit-habit-page__text">
              Opened from the habits list. The page will populate once the API responds.
            </p>
          ) : (
            <div className="edit-habit-page__summary">
              <div className="edit-habit-page__summary-item">
                <span className="edit-habit-page__summary-label">UUID</span>
                <strong className="edit-habit-page__summary-value">{loadedHabit.uuid}</strong>
              </div>
              <div className="edit-habit-page__summary-item">
                <span className="edit-habit-page__summary-label">Icon</span>
                <strong className="edit-habit-page__summary-value">{loadedHabit.icon}</strong>
              </div>
              <div className="edit-habit-page__summary-item">
                <span className="edit-habit-page__summary-label">Name</span>
                <strong className="edit-habit-page__summary-value">{loadedHabit.name}</strong>
              </div>
            </div>
          )}

          <Link className="edit-habit-page__back" to="/habits">
            Back to habits
          </Link>
        </article>
      </section>
    </AdaptiveLayout>
  );
}
