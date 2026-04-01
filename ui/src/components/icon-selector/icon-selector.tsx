import {Icon8Names} from "../../enum/icon8";
import "./icon-selector.sass";

type IconSelectorProps = {
  label: string;
  value: number;
  onChange: (value: number) => void;
};

const groups = [
  {title: "Activity", rangeStart: 100, rangeEnd: 199},
  {title: "Places", rangeStart: 200, rangeEnd: 299},
  {title: "Nature", rangeStart: 300, rangeEnd: 399},
  {title: "Health", rangeStart: 400, rangeEnd: 499},
  {title: "Food", rangeStart: 500, rangeEnd: 599},
  {title: "Animals", rangeStart: 600, rangeEnd: 699},
  {title: "Restrictions", rangeStart: 700, rangeEnd: 799},
];

function titleFromFileName(fileName: string): string {
  return fileName
    .replace(".png", "")
    .split("-")
    .map((segment) => segment.charAt(0).toUpperCase() + segment.slice(1))
    .join(" ");
}

export function IconSelector({label, value, onChange}: IconSelectorProps) {
  const iconIds = Object.keys(Icon8Names)
    .map((id) => Number(id))
    .sort((left, right) => left - right);

  return (
    <section className="icon-selector" aria-label={label}>
      <span className="icon-selector__label">{label}</span>

      <div className="icon-selector__groups">
        {groups.map((group) => {
          const groupIconIds = iconIds.filter(
            (id) => id >= group.rangeStart && id <= group.rangeEnd,
          );

          return (
            <details className="icon-selector__group" key={group.title} open>
              <summary className="icon-selector__summary">
                <span className="icon-selector__group-title">{group.title}</span>
                <span className="icon-selector__group-count">{groupIconIds.length}</span>
              </summary>

              <div className="icon-selector__grid">
                {groupIconIds.map((iconId) => {
                  const fileName = Icon8Names[iconId];
                  const title = titleFromFileName(fileName);

                  return (
                    <button
                      aria-label={title}
                      aria-pressed={value === iconId}
                      className={`icon-selector__item${value === iconId ? " icon-selector__item--selected" : ""}`}
                      key={iconId}
                      title={title}
                      type="button"
                      onClick={() => onChange(iconId)}
                    >
                      <img
                        alt={title}
                        className="icon-selector__image"
                        height="56"
                        loading="lazy"
                        src={`/img/icons/${fileName}`}
                        width="56"
                      />
                    </button>
                  );
                })}
              </div>
            </details>
          );
        })}
      </div>
    </section>
  );
}
