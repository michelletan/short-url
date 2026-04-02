import { NextResponse } from "next/server";
import { withAuth, AuthedRequest } from "@/lib/auth";
import { backendFetch } from "@/lib/api";

type Link = {
  id: number;
  slug: string;
  url: string;
  created_at: string;
  click_count: number;
};

export const GET = withAuth(async (req: AuthedRequest) => {
  try {
    const data = await backendFetch<Link[]>("/api/links", { token: req.token });
    return NextResponse.json(data);
  } catch {
    return NextResponse.json({ error: "Failed to fetch links" }, { status: 500 });
  }
});

export const POST = withAuth(async (req: AuthedRequest) => {
  const body = await req.json();
  if (!body.url) {
    return NextResponse.json({ error: "url is required" }, { status: 400 });
  }

  try {
    const data = await backendFetch<Link>("/api/links", {
      method: "POST",
      token: req.token,
      body,
    });
    return NextResponse.json(data, { status: 201 });
  } catch {
    return NextResponse.json({ error: "Failed to create link" }, { status: 500 });
  }
});