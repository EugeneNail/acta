import { DomainPreview } from "../../components/domain-preview/domain-preview";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./habits-page.sass";

export function HabitsPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Habits"
      title="Habit management supports the daily entry instead of replacing it."
      description="The existing journal backend already exposes habit CRUD. This page reserves the future UI for creating and maintaining the user's repeating daily goals."
    >
      <section className="habits-page__grid">
        <DomainPreview
          eyebrow="Catalog"
          title="Habit list placeholder"
          description="A future list of active habits, each with icon, name, and ordering controls."
        />
        <DomainPreview
          eyebrow="Editor"
          accent="orange"
          title="Habit editor placeholder"
          description="A reserved area for creating and updating the habit definition used in daily entries."
        />
      </section>
    </AdaptiveLayout>
  );
}
