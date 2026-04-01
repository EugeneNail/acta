import axios from "axios";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import type { HabitDto } from "../../dto/habit";
import { Icon8Names } from "../../enum/icon8";
import { httpClient } from "../../infrastructure/http-client/http-client";
import "./habits-page.sass";

const habitLimit = 20;

export function HabitsPage() {
  const [habits, setHabits] = useState<HabitDto[]>([]);
  const [requestError, setRequestError] = useState("");
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    async function loadHabits() {
      setIsLoading(true);
      setRequestError("");

      try {
        const response = await httpClient.get<HabitDto[]>("/api/v1/journal/habits");

        if (isMounted) {
          setHabits(response.data);
        }
      } catch (error) {
        if (!isMounted) {
          return;
        }

        if (axios.isAxiosError(error) && typeof error.response?.data === "object" && error.response?.data !== null) {
          const responseData = error.response.data as Record<string, unknown>;

          if ("message" in responseData && typeof responseData.message === "string") {
            setRequestError(responseData.message);
          } else {
            setRequestError("Failed to load habits.");
          }
        } else {
          setRequestError("Failed to load habits.");
        }
      } finally {
        if (isMounted) {
          setIsLoading(false);
        }
      }
    }

    void loadHabits();

    return () => {
      isMounted = false;
    };
  }, []);

  return (
    <AdaptiveLayout
      eyebrow="Acta / Habits"
      title="Habits define the recurring goals that can later be marked inside daily entries."
      description="Review your active habits, jump to editing, and create a new one when you need another recurring target."
    >
      <section className="habits-page__summary">
        <article className="habits-page__summary-card">
          <span className="habits-page__eyebrow">Habit capacity</span>
          <div className="habits-page__summary-line">
            <strong className="habits-page__count">{habits.length}</strong>
            <span className="habits-page__limit">/ {habitLimit}</span>
          </div>
          <p className="habits-page__text">
            Active habits currently available for daily tracking.
          </p>
          <p className="habits-page__note">
            TODO: Move habit limit retrieval to an API route.
          </p>
        </article>
      </section>

      {requestError !== "" && <div className="habits-page__request-error">{requestError}</div>}

      <section className="habits-page__grid">
        {habits.length < habitLimit && (
          <Link className="habits-page__create-card" to="/habits/create">
            <span className="habits-page__create-mark">+</span>
            <strong className="habits-page__create-title">Add habit</strong>
            <span className="habits-page__create-text">
              Create another recurring goal.
            </span>
          </Link>
        )}

        {isLoading ? (
          <article className="habits-page__card habits-page__card--state">
            <span className="habits-page__text">Loading habits...</span>
          </article>
        ) : (
          habits.map((habit) => (
            <Link className="habits-page__card habits-page__card--habit" key={habit.uuid} to={`/habits/${habit.uuid}`}>
              <div className="habits-page__icon-shell">
                <img
                  alt={habit.name}
                  className="habits-page__icon"
                  height="72"
                  src={`/img/icons/${Icon8Names[habit.icon]}`}
                  width="72"
                />
              </div>
              <div className="habits-page__content">
                <h2 className="habits-page__title">{habit.name}</h2>
                <span className="habits-page__uuid">{habit.uuid}</span>
              </div>
            </Link>
          ))
        )}

        {!isLoading && habits.length === 0 && requestError === "" && (
          <article className="habits-page__card habits-page__card--state">
            <span className="habits-page__eyebrow">No habits yet</span>
            <p className="habits-page__text">
              Create your first habit to start building the daily journal flow.
            </p>
          </article>
        )}
      </section>
    </AdaptiveLayout>
  );
}
