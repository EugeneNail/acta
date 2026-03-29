import { DomainPreview } from "../../components/domain-preview/domain-preview";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./gallery-page.sass";

export function GalleryPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Photos"
      title="Photos belong to the same daily memory stream as the journal entry."
      description="This page keeps space for upload, review, and attachment flows that will later connect each image to a specific day."
    >
      <section className="gallery-page__grid">
        <DomainPreview
          eyebrow="Uploads"
          accent="sun"
          title="Photo upload placeholder"
          description="A future drag-and-drop or camera entry point for the day."
        />
        <DomainPreview
          eyebrow="Timeline"
          accent="sage"
          title="Photo strip placeholder"
          description="A future memory strip arranged by entry date rather than by abstract media folders."
        />
        <DomainPreview
          eyebrow="Metadata"
          accent="orange"
          title="Caption and notes placeholder"
          description="A reserved place for captions and lightweight context bound to the same daily record."
        />
      </section>
    </AdaptiveLayout>
  );
}
