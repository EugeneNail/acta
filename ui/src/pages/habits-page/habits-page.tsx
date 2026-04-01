import { Link } from "react-router-dom";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import { DomainPreview } from "../../components/domain-preview/domain-preview";
import "./habits-page.sass";

export function HabitsPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Habits"
      title="Habit management supports the daily entry instead of replacing it."
      description="This page reserves the future habit catalog, while creation lives on a dedicated route."
    >
      <section className="habits-page__grid">
        <DomainPreview
          eyebrow="Catalog"
          title="Habit list placeholder"
          description="A future list of active habits, each with icon, name, and ordering controls."
        />
        <article className="habits-page__card">
          <span className="habits-page__eyebrow">Actions</span>
          <h2 className="habits-page__title">Create a new habit on a dedicated page.</h2>
          <p className="habits-page__text">
            Creation is intentionally separated from the catalog so the page can
            stay focused on list and edit workflows later.
          </p>
          <Link className="habits-page__link" to="/habits/create">
            Open create habit page
          </Link>
        </article>
      </section>
    </AdaptiveLayout>
  );
}
