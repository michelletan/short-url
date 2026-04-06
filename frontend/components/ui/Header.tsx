"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import Brand from "./Brand";

type User = {
  username: string;
};

type HeaderProps = {
  /** Skip fetching /api/me and use controlled mode */
  user?: User | null;
  /** Show logout button (dashboard mode) */
  showLogout?: boolean;
  onLogout?: () => void;
};

export function Header({
  user: controlledUser,
  showLogout = false,
  onLogout,
}: HeaderProps) {
  const [user, setUser] = useState<User | null>(controlledUser ?? null);
  const [checked, setChecked] = useState(!!controlledUser);

  // Only fetch if user is NOT provided
  useEffect(() => {
    if (controlledUser !== undefined) return;

    fetch("/api/me")
      .then((res) => (res.ok ? res.json() : null))
      .then((data) => {
        setUser(data ?? null);
        setChecked(true);
      })
      .catch(() => setChecked(true));
  }, [controlledUser]);

  return (
    <header className="w-full border-b border-[#1a1a1a]">
      <div className="max-w-4xl mx-auto px-6 h-14 flex items-center justify-between">
        <Brand />

        {checked && (
          <nav className="flex items-center gap-6">
            {user ? (
              <>
                <Link
                  href="/dashboard"
                  className="font-mono text-xs text-[#555] hover:text-[#e8ff47]"
                >
                  <span className="text-[#333] mr-1">@</span>
                  {user.username}
                </Link>

                {showLogout && onLogout && (
                  <button
                    onClick={onLogout}
                    className="font-mono text-[11px] tracking-widest text-[#444] hover:text-[#e8ff47] uppercase transition-colors duration-200"
                  >
                    Sign out
                  </button>
                )}
              </>
            ) : (
              <>
                <Link
                  href="/login"
                  className="font-mono text-xs tracking-widest text-[#444] hover:text-[#e8ff47] uppercase"
                >
                  Login
                </Link>
                <Link
                  href="/register"
                  className="font-mono text-xs tracking-widest uppercase px-4 py-2 border border-[#2a2a2a] text-[#666] hover:border-[#e8ff47] hover:text-[#e8ff47] rounded-sm"
                >
                  Register
                </Link>
              </>
            )}
          </nav>
        )}
      </div>
    </header>
  );
}