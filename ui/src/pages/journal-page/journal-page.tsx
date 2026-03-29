import { PageLayout } from "../../layouts/page-layout/page-layout";
import "./journal-page.sass";

export function JournalPage() {
  return (
    <PageLayout
      eyebrow="Acta / Journal"
      title="Journal routes stay separate from the auth surface."
      description="The page is still a placeholder, but it now sits next to real authentication pages rather than temporary router stubs."
    >
      <section className="journalStub">
        <p>Journal UI will be built here after the habit screens are connected.</p>
      </section>
    </PageLayout>
  );
}
