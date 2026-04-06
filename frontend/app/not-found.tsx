import Link from "next/link";
import Brand from "@/components/ui/Brand";
import { theme } from "@/lib/theme";

export default function NotFound() {
  return (
    <div className="min-h-screen bg-[#0e0e0e] text-white flex flex-col items-center justify-center px-6">
      <div className={theme.gridBg.className} style={theme.gridBg.style} />

      <div className="relative text-center">
        {/* Big 404 */}
        <p className="font-mono text-[10rem] font-light leading-none text-[#1a1a1a] select-none">
          404
        </p>

        {/* Overlaid message */}
        <div className="-mt-10">
          <div className="flex items-center justify-center gap-3 mb-4">
            <div className="w-6 h-px bg-[#e8ff47]" />
            <span className="font-mono text-[11px] tracking-[0.3em] text-[#e8ff47] uppercase">
              Link not found
            </span>
            <div className="w-6 h-px bg-[#e8ff47]" />
          </div>

          <p className="text-[#444] text-sm font-light mb-10">
            This link doesn't exist or may have been removed.
          </p>

          <Link
            href="/"
            className="inline-block font-mono text-xs tracking-widest uppercase px-7 py-3.5 bg-[#e8ff47] text-[#0e0e0e] rounded-sm hover:bg-white transition-colors duration-200"
          >
            Back to home →
          </Link>
        </div>
      </div>

      <Brand />
    </div>
  );
}