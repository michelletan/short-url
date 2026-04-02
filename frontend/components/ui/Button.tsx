import { theme } from "@/lib/theme";

type ButtonVariant = "primary" | "ghost";

type ButtonProps = {
  variant?: ButtonVariant;
} & React.ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({
  variant = "primary",
  className = "",
  children,
  ...props
}: ButtonProps) {
  return (
    <button
      className={[theme.button[variant], className].filter(Boolean).join(" ")}
      {...props}
    >
      {children}
    </button>
  );
}
