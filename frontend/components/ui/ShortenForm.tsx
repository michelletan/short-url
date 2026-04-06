type Props = {
  url: string;
  setUrl: (val: string) => void;
  onSubmit: (e: React.FormEvent) => void;
  submitting: boolean;
  error: string;
};

export default function ShortenForm({
  url,
  setUrl,
  onSubmit,
  submitting,
  error,
}: Props) {
  return (
    <section className="mb-16">
      <h1 className="text-2xl font-light tracking-tight mb-6">
        Shorten a URL
      </h1>

      <form onSubmit={onSubmit} className="flex gap-3">
        <input
          type="url"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          required
          placeholder="https://your-very-long-url.com/goes/here"
          className="flex-1 bg-[#161616] border border-[#2a2a2a] rounded-sm px-4 py-3 text-sm text-white placeholder-[#444] focus:outline-none focus:border-[#e8ff47]"
        />

        <button
          type="submit"
          disabled={submitting}
          className="bg-[#e8ff47] text-[#0e0e0e] font-mono text-xs tracking-widest uppercase px-6 py-3 rounded-sm disabled:opacity-40"
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
  );
}