import "./error-list.sass";

type Props = {
  errors: Record<string, string>;
};

export function ErrorList({ errors }: Props) {
  const entries = Object.entries(errors);

  if (entries.length === 0) {
    return null;
  }

  return (
    <ul className="errorList">
      {entries.map(([field, message]) => (
        <li key={`${field}-${message}`}>
          <span>{field}</span>
          <strong>{message}</strong>
        </li>
      ))}
    </ul>
  );
}
