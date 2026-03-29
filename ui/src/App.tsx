export default function App() {
  return (
    <main className="shell">
      <section className="hero">
        <p className="eyebrow">Acta / UI</p>
        <h1>Build daily momentum with a calmer interface.</h1>
        <p className="lede">
          The UI service is ready. Next steps are wiring real auth state,
          journal screens, and API calls through the shared proxy.
        </p>
        <div className="actions">
          <a className="primaryAction" href="/api/v1/auth/login">
            Auth API
          </a>
          <a className="secondaryAction" href="/api/v1/journal/habits">
            Journal API
          </a>
        </div>
      </section>

      <section className="panelGrid">
        <article className="panel highlight">
          <span className="panelTag">Today</span>
          <h2>Journal-first workspace</h2>
          <p>
            A front-end shell for the Acta platform, designed to sit behind the
            shared proxy and evolve with the microservices.
          </p>
        </article>

        <article className="panel">
          <span className="panelTag">Auth</span>
          <h2>Token-aware entrypoint</h2>
          <p>
            Ready for login, refresh flows, and route guards once state
            management is added.
          </p>
        </article>

        <article className="panel">
          <span className="panelTag">Journal</span>
          <h2>Habit surface</h2>
          <p>
            Prepared to consume the existing habit routes and expand into the
            full journal experience.
          </p>
        </article>
      </section>
    </main>
  );
}
