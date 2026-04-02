"use client";

import { useState, useEffect, useCallback } from "react";
import { useRouter } from "next/navigation";

type Link = {
  id: number;
  url: string;
  slug: string;
  created_at: string;
  click_count: number;
};

export default function DashboardPage() {
  const router = useRouter();
  const [links, setLinks] = useState<Link[]>([]);
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");
  const [copiedSlug, setCopiedSlug] = useState<string | null>(null);

  const fetchLinks = useCallback(async () => {
    try {
      const res = await fetch("/api/links");
      if (res.status === 401) {
        router.push("/login");
        return;
      }
      const data = await res.json();
      setLinks(data.links ?? []);
    } catch {
      setError("Failed to load links");
    } finally {
      setLoading(false);
    }
  }, [router]);

  useEffect(() => {
    fetchLinks();
  }, [fetchLinks]);

  async function handleShorten(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    setSubmitting(true);

    try {
      const res = await fetch("/api/links", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url }),
      });

      if (!res.ok) {
        const data = await res.json();
        setError(data.error || "Failed to shorten URL");
        return;
      }

      setUrl("");
      fetchLinks();
    } catch {
      setError("Something went wrong. Please try again.");
    } finally {
      setSubmitting(false);
    }
  }

  async function handleLogout() {
    await fetch("/api/auth/logout", { method: "POST" });
    router.push("/login");
    router.refresh();
  }

  function handleCopy(slug: string) {
    const shortUrl = `${window.location.origin}/${slug}`;
    navigator.clipboard.writeText(shortUrl);
    setCopiedSlug(slug);
    setTimeout(() => setCopiedSlug(null), 2000);
  }

  return (
    <div className="min-h-screen bg-[#0e0e0e] text-white">
      {/* Subtle grid background */}
      <div
        className="fixed inset-0 opacity-[0.03] pointer-events-none"
        style={{
          backgroundImage:
            "linear-gradient(#fff 1px, transparent 1px), linear-gradient(90deg, #fff 1px, transparent 1px)",
          backgroundSize: "64px 64px",
        }}
      />

      <div className="relative max-w-3xl mx-auto px-6 py-12">
        {/* Header */}
        <header className="flex items-center justify-between mb-16">
          <span className="font-mono text-xs tracking-[0.3em] text-[#666] uppercase">
            snip
          </span>
          <button
            onClick={handleLogout}
            className="font-mono text-[11px] tracking-widest text-[#444] hover:text-[#e8ff47] uppercase transition-colors duration-200"
          >
            Sign out
          </button>
        </header>

        {/* Shorten form */}
        <section className="mb-16">
          <h1 className="text-2xl font-light tracking-tight mb-6">
            Shorten a URL
          </h1>
          <form onSubmit={handleShorten} className="flex gap-3">
            <input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              required
              placeholder="https://your-very-long-url.com/goes/here"
              className="flex-1 bg-[#161616] border border-[#2a2a2a] rounded-sm px-4 py-3 text-sm text-white placeholder-[#444] focus:outline-none focus:border-[#e8ff47] transition-colors duration-200"
            />
            <button
              type="submit"
              disabled={submitting}
              className="bg-[#e8ff47] text-[#0e0e0e] font-mono text-xs tracking-widest uppercase px-6 py-3 rounded-sm hover:bg-white transition-colors duration-200 disabled:opacity-40 disabled:cursor-not-allowed whitespace-nowrap"
            >
              {submitting ? "..." : "Snip →"}
            </button>
          </form>

          {error && (
            <p className="mt-3 font-mono text-xs text-red-400 border border-red-400/20 bg-red-400/5 px-3 py-2 rounded-sm">
              {error}
            </p>
          )}
        </section>

        {/* Divider */}
        <div className="border-t border-[#1e1e1e] mb-10" />

        {/* Links list */}
        <section>
          <div className="flex items-center justify-between mb-6">
            <h2 className="font-mono text-[11px] tracking-widest text-[#555] uppercase">
              Your links
            </h2>
            <span className="font-mono text-[11px] text-[#333]">
              {links.length} total
            </span>
          </div>

          {loading ? (
            <div className="space-y-3">
              {[...Array(3)].map((_, i) => (
                <div
                  key={i}
                  className="h-16 bg-[#161616] rounded-sm animate-pulse"
                />
              ))}
            </div>
          ) : links.length === 0 ? (
            <p className="font-mono text-xs text-[#333] py-8 text-center border border-dashed border-[#1e1e1e] rounded-sm">
              No links yet. Snip something above.
            </p>
          ) : (
            <ul className="space-y-2">
              {links.map((link) => (
                <li
                  key={link.id}
                  className="group flex items-center justify-between bg-[#131313] border border-[#1e1e1e] hover:border-[#2a2a2a] rounded-sm px-5 py-4 transition-colors duration-200"
                >
                  <div className="min-w-0 flex-1 mr-6">
                    {/* Short URL */}
                    <p className="font-mono text-sm text-[#e8ff47] truncate">
                      {window.location.host}/{link.slug}
                    </p>
                    {/* Original URL */}
                    <p className="text-xs text-[#444] truncate mt-0.5">
                      {link.url}
                    </p>
                  </div>

                  <div className="flex items-center gap-5 shrink-0">
                    {/* Click count */}
                    <div className="text-right">
                      <p className="font-mono text-sm text-white">
                        {link.click_count}
                      </p>
                      <p className="font-mono text-[10px] text-[#444] uppercase tracking-wider">
                        clicks
                      </p>
                    </div>

                    {/* Copy button */}
                    <button
                      onClick={() => handleCopy(link.slug)}
                      className="font-mono text-[11px] tracking-widest uppercase text-[#444] hover:text-[#e8ff47] transition-colors duration-200 w-12 text-right"
                    >
                      {copiedSlug === link.slug ? "✓" : "Copy"}
                    </button>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </section>
      </div>
    </div>
  );
}