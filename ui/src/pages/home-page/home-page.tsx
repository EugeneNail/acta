import { PageLayout } from "../../layouts/page-layout/page-layout";
import "./home-page.sass";

export function HomePage() {
  return (
    <PageLayout
      eyebrow="Acta / Home"
      title="A quiet interface over a growing microservice stack."
      description="The UI now has dedicated auth routes connected to the existing signup and login flows through the shared proxy."
    >
      <section className="panelGrid">
        <article className="panel highlight">
          <p className="panelTag">Authentication</p>
          <h2>Register and sign in against the live auth service.</h2>
          <p>
            The pages use the current backend contract and surface transport and
            application validation errors directly in the UI.
          </p>
        </article>
        <article className="panel">
          <p className="panelTag">Journal</p>
          <h2>Journal stays behind authenticated API routes.</h2>
          <p>
            The next step is to use the issued access token for the journal
            habit screens.
          </p>
        </article>
        <article className="panel">
          <p className="panelTag">Routing</p>
          <h2>Client-side routes survive proxy reloads.</h2>
          <p>
            The UI router is served through the proxy, while API paths stay
            isolated under <code>/api</code>.
          </p>
        </article>
      </section>
    </PageLayout>
  );
}
