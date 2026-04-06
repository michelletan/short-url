"use client";

import { useState, useEffect, useCallback } from "react";
import { useRouter } from "next/navigation";

import { Header } from "@/components/ui/Header";
import ShortenForm from "@/components/ui/ShortenForm";
import LinksList from "@/components/ui/LinksList";

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
      <div className="relative max-w-3xl mx-auto px-6 py-12">
        <Header
          showLogout
          onLogout={handleLogout}
        />

        <ShortenForm
          url={url}
          setUrl={setUrl}
          onSubmit={handleShorten}
          submitting={submitting}
          error={error}
        />

        <div className="border-t border-[#1e1e1e] mb-10" />

        <LinksList
          links={links}
          loading={loading}
          copiedSlug={copiedSlug}
          onCopy={handleCopy}
        />
      </div>
    </div>
  );
}