type Link = {
  id: number;
  url: string;
  slug: string;
  click_count: number;
};

type Props = {
  link: Link;
  copiedSlug: string | null;
  onCopy: (slug: string) => void;
};

export default function LinkItem({ link, copiedSlug, onCopy }: Props) {
  return (
    <li className="group flex items-center justify-between bg-[#131313] border border-[#1e1e1e] rounded-sm px-5 py-4">
      <div className="min-w-0 flex-1 mr-6">
        <p className="font-mono text-sm text-[#e8ff47] truncate">
          {window.location.host}/{link.slug}
        </p>

        <p className="text-xs text-[#444] truncate mt-0.5">
          {link.url}
        </p>
      </div>

      <div className="flex items-center gap-5">
        <div className="text-right">
          <p className="font-mono text-sm text-white">
            {link.click_count}
          </p>
          <p className="font-mono text-[10px] text-[#444] uppercase">
            clicks
          </p>
        </div>

        <button
          onClick={() => onCopy(link.slug)}
          className="font-mono text-[11px] uppercase text-[#444] hover:text-[#e8ff47]"
        >
          {copiedSlug === link.slug ? "✓" : "Copy"}
        </button>
      </div>
    </li>
  );
}