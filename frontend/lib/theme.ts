/**
 * Central theme configuration.
 * All component styles are derived from this file.
 */

export const theme = {
  // ─── Colours ────────────────────────────────────────────────
  colors: {
    bg: {
      page: "#0e0e0e",
      surface: "#131313",
      input: "#161616",
    },
    border: {
      default: "#2a2a2a",
      subtle: "#1e1e1e",
      focus: "#e8ff47",
    },
    text: {
      primary: "#ffffff",
      muted: "#666666",
      faint: "#444444",
      placeholder: "#444444",
      label: "#555555",
    },
    accent: {
      primary: "#e8ff47",    // yellow-green — main CTA
      hover: "#ffffff",      // accent button hover
      text: "#0e0e0e",       // text on accent bg
    },
    status: {
      errorText: "#f87171",  // red-400
      errorBorder: "rgba(248,113,113,0.2)",
      errorBg: "rgba(248,113,113,0.05)",
    },
  },

  // ─── Typography ─────────────────────────────────────────────
  font: {
    sans: "font-['Sora',sans-serif]",
    mono: "font-mono",
  },

  // ─── Radii ──────────────────────────────────────────────────
  radius: {
    sm: "rounded-sm",
    md: "rounded-md",
  },

  // ─── Reusable class strings ──────────────────────────────────
  input: {
    base: "w-full bg-[#161616] border border-[#2a2a2a] rounded-sm px-4 py-3 text-sm text-white placeholder-[#444] focus:outline-none focus:border-[#e8ff47] transition-colors duration-200",
    error: "border-red-400 focus:border-red-400",
  },

  button: {
    primary:
      "w-full bg-[#e8ff47] text-[#0e0e0e] font-mono text-xs tracking-widest uppercase px-4 py-3.5 rounded-sm hover:bg-white transition-colors duration-200 disabled:opacity-40 disabled:cursor-not-allowed",
    ghost:
      "font-mono text-[11px] tracking-widest text-[#444] hover:text-[#e8ff47] uppercase transition-colors duration-200",
  },

  label: "block font-mono text-[11px] tracking-widest text-[#555] uppercase",

  errorMessage:
    "font-mono text-xs text-red-400 border border-red-400/20 bg-red-400/5 px-3 py-2 rounded-sm",

  // ─── Background decoration ───────────────────────────────────
  gridBg: {
    style: {
      backgroundImage:
        "linear-gradient(#fff 1px, transparent 1px), linear-gradient(90deg, #fff 1px, transparent 1px)",
      backgroundSize: "64px 64px",
    },
    className: "fixed inset-0 opacity-[0.03] pointer-events-none",
  },
} as const;
