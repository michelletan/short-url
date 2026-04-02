import Link from "next/link";
import { Header } from "@/components/ui/Header";
import { theme } from "@/lib/theme";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default async function LandingPage() {
  const cookieStore = await cookies();
  const session = cookieStore.get("session");

  if (session) redirect("/dashboard");

  return (
    <div className="min-h-screen bg-[#0e0e0e] text-white flex flex-col">
      <div className={theme.gridBg.className} style={theme.gridBg.style} />

      <Header />

      <main className="relative flex-1 flex flex-col">
        <section className="flex-1 flex flex-col justify-center max-w-4xl mx-auto px-6 py-24">
          <div className="flex items-center gap-3 mb-8">
            <div className="w-8 h-px bg-[#e8ff47]" />
            <span className="font-mono text-[11px] tracking-[0.3em] text-[#e8ff47] uppercase">
              Short URL
            </span>
          </div>

          <h1 className="text-5xl sm:text-6xl font-light tracking-tight leading-[1.1] mb-6">
            Long URLs are
            <br />
            <span className="text-[#e8ff47]">ugly.</span> Fix that.
          </h1>

          <p className="text-[#555] text-lg font-light leading-relaxed max-w-md mb-12">
            Short URL turns any URL into something clean and shareable.
            Track clicks, manage your links — all in one place.
            <br />
            <span className="text-[#3a3a3a] text-sm mt-2 block">
              Account required to create and manage links.
            </span>
          </p>

          <div className="flex items-center gap-4">
            <Link
              href="/register"
              className="inline-block bg-[#e8ff47] text-[#0e0e0e] font-mono text-xs tracking-widest uppercase px-7 py-3.5 rounded-sm hover:bg-white transition-colors duration-200"
            >
              Get started →
            </Link>
            <Link
              href="/login"
              className="font-mono text-xs tracking-widest text-[#444] hover:text-[#e8ff47] uppercase transition-colors duration-200"
            >
              Sign in
            </Link>
          </div>
        </section>

        <section className="border-t border-[#1a1a1a]">
          <div className="max-w-4xl mx-auto px-6 py-16 grid grid-cols-1 sm:grid-cols-3 gap-10">
            {[
              {
                label: "Shorten",
                description:
                  "Paste any URL and get a clean short link in one click.",
              },
              {
                label: "Track",
                description:
                  "See how many times each link has been clicked, in real time.",
              },
              {
                label: "Manage",
                description:
                  "All your links in one dashboard. Copy, share, or review them anytime.",
              },
            ].map(({ label, description }) => (
              <div key={label}>
                <p className="font-mono text-[11px] tracking-widest text-[#e8ff47] uppercase mb-3">
                  {label}
                </p>
                <p className="text-sm text-[#555] font-light leading-relaxed">
                  {description}
                </p>
              </div>
            ))}
          </div>
        </section>

        <footer className="border-t border-[#1a1a1a]">
          <div className="max-w-4xl mx-auto px-6 py-6 flex items-center justify-between">
            <span className="font-mono text-[11px] tracking-[0.3em] text-[#333] uppercase">
              snip
            </span>
            <span className="font-mono text-[11px] text-[#2a2a2a]">
              Built with Go + Next.js
            </span>
          </div>
        </footer>
      </main>
    </div>
  );
}