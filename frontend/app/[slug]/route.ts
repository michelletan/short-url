import { NextRequest, NextResponse } from "next/server";

const REDIRECT_BASE_URL = process.env.REDIRECT_BASE_URL ?? "http://localhost:8080";

export async function GET(
  req: NextRequest,
  { params }: { params: Promise<{ slug: string }> }
) {
  const { slug } = await params;

  const ip = req.headers.get("x-forwarded-for")?.split(",")[0].trim() ??
    req.headers.get("x-real-ip") ??
    "";

  const res = await fetch(`${REDIRECT_BASE_URL}/${slug}`, {
    redirect: "manual",
    headers: {
      "User-Agent": req.headers.get("user-agent") ?? "",
      "X-Forwarded-For": ip,
      "Referer": req.headers.get("referer") ?? "",
    },
  });

  if (res.status === 301 || res.status === 302 || res.status === 307) {
    const location = res.headers.get("location");
    if (location) {
      return NextResponse.redirect(location, { status: res.status });
    }
  }

  if (res.status === 404) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  return NextResponse.redirect(new URL("/", req.url));
}