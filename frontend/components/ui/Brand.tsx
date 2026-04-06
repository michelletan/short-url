import Link from "next/link";
import React from "react";

interface BrandProps {
  href?: string;
}

const Brand: React.FC<BrandProps> = ({ href = "/" }) => {
  return (
    <Link
      href={href}
      className="font-mono text-xs tracking-[0.3em] text-[#666] uppercase hover:text-[#e8ff47] transition-colors duration-200"
    >
      Short URL
    </Link>
  );
};

export default Brand;