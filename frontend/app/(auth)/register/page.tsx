"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { Input } from "@/components/ui/Input";
import { Button } from "@/components/ui/Button";
import { theme } from "@/lib/theme";

type FormErrors = {
  email?: string;
  username?: string;
  password?: string;
};

function validate(email: string, username: string, password: string): FormErrors {
  const errors: FormErrors = {};

  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    errors.email = "Enter a valid email address";
  }
  if (username.length < 8 || username.length > 30) {
    errors.username = "Username must be 8–30 characters";
  }
  if (password.length < 8 || password.length > 30) {
    errors.password = "Password must be 8–30 characters";
  }

  return errors;
}

export default function RegisterPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({});
  const [serverError, setServerError] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setServerError("");

    const errors = validate(email, username, password);
    if (Object.keys(errors).length > 0) {
      setFieldErrors(errors);
      return;
    }
    setFieldErrors({});
    setLoading(true);

    try {
      const res = await fetch("/api/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, username, password }),
      });

      if (!res.ok) {
        const data = await res.json();
        setServerError(data.error || "Registration failed");
        return;
      }

      // Auto-login after register
      const loginRes = await fetch("/api/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      router.push(loginRes.ok ? "/dashboard" : "/login");
      router.refresh();
    } catch {
      setServerError("Something went wrong. Please try again.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="min-h-screen bg-[#0e0e0e] flex items-center justify-center px-4">
      <div className={theme.gridBg.className} style={theme.gridBg.style} />

      <div className="relative w-full max-w-sm">
        <div className="mb-12">
          <span className="font-mono text-xs tracking-[0.3em] text-[#666] uppercase">
            snip
          </span>
          <h1 className="mt-3 text-3xl font-light text-white tracking-tight">
            Create account.
          </h1>
        </div>

        <form onSubmit={handleSubmit} className="space-y-5">
          <Input
            label="Email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            autoComplete="email"
            placeholder="you@example.com"
            error={fieldErrors.email}
          />
          <Input
            label="Username"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            autoComplete="username"
            placeholder="8–30 characters"
            minLength={8}
            maxLength={30}
            error={fieldErrors.username}
          />
          <Input
            label="Password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            autoComplete="new-password"
            placeholder="8–30 characters"
            minLength={8}
            maxLength={30}
            error={fieldErrors.password}
          />

          {serverError && (
            <p className={theme.errorMessage}>{serverError}</p>
          )}

          <Button type="submit" disabled={loading} className="mt-2">
            {loading ? "Creating account..." : "Create account →"}
          </Button>
        </form>

        <p className="mt-8 font-mono text-xs text-[#444]">
          Already have an account?{" "}
          <Link
            href="/login"
            className="text-[#666] hover:text-[#e8ff47] transition-colors duration-200"
          >
            Sign in
          </Link>
        </p>
      </div>
    </div>
  );
}
