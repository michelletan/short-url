import LinkItem from "./LinkItem";

type Link = {
  id: number;
  url: string;
  slug: string;
  click_count: number;
};

type Props = {
  links: Link[];
  loading: boolean;
  copiedSlug: string | null;
  onCopy: (slug: string) => void;
};

export default function LinksList({
  links,
  loading,
  copiedSlug,
  onCopy,
}: Props) {
  if (loading) {
    return (
      <div className="space-y-3">
        {[...Array(3)].map((_, i) => (
          <div key={i} className="h-16 bg-[#161616] rounded-sm animate-pulse" />
        ))}
      </div>
    );
  }

  if (links.length === 0) {
    return (
      <p className="font-mono text-xs text-[#333] py-8 text-center border border-dashed border-[#1e1e1e] rounded-sm">
        No links yet. Snip something above.
      </p>
    );
  }

  return (
    <ul className="space-y-2">
      {links.map((link) => (
        <LinkItem
          key={link.id}
          link={link}
          copiedSlug={copiedSlug}
          onCopy={onCopy}
        />
      ))}
    </ul>
  );
}