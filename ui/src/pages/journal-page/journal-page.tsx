import { DomainPreview } from "../../components/domain-preview/domain-preview";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./journal-page.sass";

export function JournalPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Journal Entry"
      title="The daily entry page will become the center of the product."
      description="This screen is reserved for the future writing flow: date context, mood, habit completion, and attached photo blocks inside one daily record."
    >
      <section className="journal-page__grid">
        <DomainPreview
          eyebrow="Writing"
          title="Main diary editor placeholder"
          description="A future text surface for reflection, observations, and day summaries."
        />
        <DomainPreview
          eyebrow="Mood"
          accent="sun"
          title="Mood selector placeholder"
          description="A simplified visual control for the feeling of the day."
        />
      </section>
    </AdaptiveLayout>
  );
}
