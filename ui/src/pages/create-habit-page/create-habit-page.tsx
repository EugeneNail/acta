import type { FormEvent } from "react";
import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import {IconSelector} from "../../components/icon-selector/icon-selector";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import type {
  CreateHabitDto,
  HabitDto,
  ValidationErrorsDto,
} from "../../dto/habit";
import { httpClient } from "../../infrastructure/http-client/http-client";
import "./create-habit-page.sass";

export function CreateHabitPage() {
  const navigate = useNavigate();
  const [icon, setIcon] = useState(100);
  const [name, setName] = useState("");
  const [fieldErrors, setFieldErrors] = useState<ValidationErrorsDto>({});
  const [requestError, setRequestError] = useState("");
  const [createdHabit, setCreatedHabit] = useState<HabitDto | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setFieldErrors({});
    setRequestError("");
    setCreatedHabit(null);

    try {
      const payload: CreateHabitDto = {
        icon,
        name,
      };

      const response = await httpClient.post<HabitDto>("/api/v1/journal/habits", payload);
      setCreatedHabit(response.data);
      setName("");
      navigate("/habits");
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

  return (
    <AdaptiveLayout
      eyebrow="Acta / Create Habit"
      title="Create a habit that will later appear inside daily entries."
      description="This page sends the create-habit payload directly to the journal API."
    >
      <section className="create-habit-page__grid">
        <article className="create-habit-page__card create-habit-page__card--form">
          <span className="create-habit-page__eyebrow">Create habit</span>
          <h2 className="create-habit-page__title">New daily goal</h2>

          <form className="create-habit-page__form" onSubmit={onSubmit}>
            <label className="create-habit-page__field">
              <span className="create-habit-page__label">Name</span>
              <input
                className="create-habit-page__input"
                name="name"
                type="text"
                value={name}
                onChange={(event) => setName(event.target.value)}
              />
            </label>

            <IconSelector label="Icon" value={icon} onChange={setIcon} />

            <button className="create-habit-page__submit" disabled={isSubmitting} type="submit">
              {isSubmitting ? "Creating..." : "Create habit"}
            </button>
          </form>

          {Object.keys(fieldErrors).length > 0 && (
            <ul className="create-habit-page__errors">
              {Object.entries(fieldErrors).map(([field, message]) => (
                <li className="create-habit-page__error" key={`${field}-${message}`}>
                  <span className="create-habit-page__error-field">{field}</span>
                  <strong className="create-habit-page__error-message">{message}</strong>
                </li>
              ))}
            </ul>
          )}

          {requestError !== "" && (
            <div className="create-habit-page__request-error">{requestError}</div>
          )}
        </article>

        <article className="create-habit-page__card create-habit-page__card--result">
          <span className="create-habit-page__eyebrow">Result</span>
          <h2 className="create-habit-page__title">Created habit</h2>

          {createdHabit === null ? (
            <p className="create-habit-page__text">
              Submit the form to create a habit through the live journal API.
            </p>
          ) : (
            <div className="create-habit-page__summary">
              <div className="create-habit-page__summary-item">
                <span className="create-habit-page__summary-label">UUID</span>
                <strong className="create-habit-page__summary-value">{createdHabit.uuid}</strong>
              </div>
              <div className="create-habit-page__summary-item">
                <span className="create-habit-page__summary-label">Icon</span>
                <strong className="create-habit-page__summary-value">{createdHabit.icon}</strong>
              </div>
              <div className="create-habit-page__summary-item">
                <span className="create-habit-page__summary-label">Name</span>
                <strong className="create-habit-page__summary-value">{createdHabit.name}</strong>
              </div>
            </div>
          )}
        </article>
      </section>
    </AdaptiveLayout>
  );
}
