"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { theme } from "@/lib/theme";

type User = {
  username: string;
};

export function Header() {
  const [user, setUser] = useState<User | null>(null);
  const [checked, setChecked] = useState(false);

  useEffect(() => {
    fetch("/api/me")
      .then((res) => (res.ok ? res.json() : null))
      .then((data) => {
        setUser(data ?? null);
        setChecked(true);
      })
      .catch(() => setChecked(true));
  }, []);

  return (
    <header className="w-full border-b border-[#1a1a1a]">
      <div className="max-w-4xl mx-auto px-6 h-14 flex items-center justify-between">
        {/* Brand */}
        <Link
          href="/"
          className="font-mono text-xs tracking-[0.3em] text-[#666] uppercase hover:text-[#e8ff47] transition-colors duration-200"
        >
          snip
        </Link>

        {/* Nav — only render once auth state is known to avoid flash */}
        {checked && (
          <nav>
            {user ? (
              <Link
                href="/dashboard"
                className="font-mono text-xs text-[#555] hover:text-[#e8ff47] transition-colors duration-200"
              >
                <span className="text-[#333] mr-1">@</span>
                {user.username}
              </Link>
            ) : (
              <div className="flex items-center gap-6">
                <Link
                  href="/login"
                  className="font-mono text-xs tracking-widest text-[#444] hover:text-[#e8ff47] uppercase transition-colors duration-200"
                >
                  Login
                </Link>
                <Link
                  href="/register"
                  className="font-mono text-xs tracking-widest uppercase px-4 py-2 border border-[#2a2a2a] text-[#666] hover:border-[#e8ff47] hover:text-[#e8ff47] rounded-sm transition-colors duration-200"
                >
                  Register
                </Link>
              </div>
            )}
          </nav>
        )}
      </div>
    </header>
  );
}