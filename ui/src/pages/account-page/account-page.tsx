import { DomainPreview } from "../../components/domain-preview/domain-preview";
import { AdaptiveLayout } from "../../layouts/adaptive-layout/adaptive-layout";
import "./account-page.sass";

export function AccountPage() {
  return (
    <AdaptiveLayout
      eyebrow="Acta / Account"
      title="Authentication and profile controls will live here."
      description="The auth service already supports signup, login, and token refresh. This page keeps a clear future space for session and profile controls."
    >
      <section className="account-page__grid">
        <DomainPreview
          eyebrow="Session"
          title="Login and refresh placeholder"
          description="A future account surface for session state, issued tokens, and device-level access."
        />
        <DomainPreview
          eyebrow="Profile"
          accent="sun"
          title="Profile placeholder"
          description="A reserved area for user-facing settings and personal metadata."
        />
      </section>
    </AdaptiveLayout>
  );
}
