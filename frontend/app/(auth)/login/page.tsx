"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { Input } from "@/components/ui/Input";
import { Button } from "@/components/ui/Button";
import { theme } from "@/lib/theme";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      const res = await fetch("/api/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (!res.ok) {
        const data = await res.json();
        setError(data.error || "Invalid credentials");
        return;
      }

      router.push("/dashboard");
      router.refresh();
    } catch {
      setError("Something went wrong. Please try again.");
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
            Welcome back.
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
          />
          <Input
            label="Password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            autoComplete="current-password"
            placeholder="••••••••"
          />

          {error && <p className={theme.errorMessage}>{error}</p>}

          <Button type="submit" disabled={loading} className="mt-2">
            {loading ? "Signing in..." : "Sign in →"}
          </Button>
        </form>

        <p className="mt-8 font-mono text-xs text-[#444]">
          No account?{" "}
          <Link
            href="/register"
            className="text-[#666] hover:text-[#e8ff47] transition-colors duration-200"
          >
            Register
          </Link>
        </p>
      </div>
    </div>
  );
}