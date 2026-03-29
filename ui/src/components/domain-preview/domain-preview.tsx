import "./domain-preview.sass";

type Props = {
  eyebrow: string;
  title: string;
  description: string;
  accent?: "sage" | "orange" | "sun";
};

export function DomainPreview({
  eyebrow,
  title,
  description,
  accent = "sage",
}: Props) {
  return (
    <article className={`domain-preview domain-preview--${accent}`}>
      <span className="domain-preview__eyebrow">{eyebrow}</span>
      <h3 className="domain-preview__title">{title}</h3>
      <p className="domain-preview__description">{description}</p>
    </article>
  );
}
