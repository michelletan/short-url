import { redirect } from "next/navigation";

export default async function Page({
  params,
}: {
  params: { slug: string };
}) {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/${params.slug}`,
    { cache: "no-store" }
  );

  if (!res.ok) {
    return <div>Link not found</div>;
  }

  const data = await res.json();

  redirect(data.originalUrl);
}