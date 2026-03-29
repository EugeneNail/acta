import { DomainPreview } from "../../components/domain-preview/domain-preview";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./dashboard-page.sass";

export function DashboardPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Domain Overview"
      title="A daily tracker shaped around entries, goals, and photo memories."
      description="The product domain centers on a single daily journal entry, completed habits, attached photos, and the continuity of everyday reflection."
    >
      <section className="dashboard-page__grid">
        <DomainPreview
          accent="sage"
          eyebrow="Daily entry"
          title="One clear place for a date, notes, mood, and the shape of the day."
          description="This card stands in for the future entry workspace where users will write the diary text and review what happened today."
        />
        <DomainPreview
          accent="sun"
          eyebrow="Goals and habits"
          title="A repeating set of daily goals that can be checked off inside the entry."
          description="This part of the product already has a backend habit service and will later become the daily completion surface."
        />
        <DomainPreview
          accent="orange"
          eyebrow="Photo timeline"
          title="Memory snapshots connected to the same day."
          description="This placeholder reserves space for attached images that enrich a journal entry rather than living as a separate feed."
        />
      </section>
    </AdaptiveLayout>
  );
}
