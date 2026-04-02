import { theme } from "@/lib/theme";

type InputProps = {
  label: string;
  error?: string;
} & React.InputHTMLAttributes<HTMLInputElement>;

export function Input({ label, error, id, ...props }: InputProps) {
  const inputId = id ?? label.toLowerCase().replace(/\s+/g, "-");

  return (
    <div className="space-y-1">
      <label htmlFor={inputId} className={theme.label}>
        {label}
      </label>
      <input
        id={inputId}
        className={[
          theme.input.base,
          error ? theme.input.error : "",
        ]
          .filter(Boolean)
          .join(" ")}
        {...props}
      />
      {error && <p className={theme.errorMessage}>{error}</p>}
    </div>
  );
}
