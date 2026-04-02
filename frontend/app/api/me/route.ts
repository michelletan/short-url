import { NextResponse } from "next/server";
import { withAuth, AuthedRequest } from "@/lib/auth";
import { backendFetch } from "@/lib/api";

type MeResponse = {
  email: string;
  username: string;
};

export const GET = withAuth(async (req: AuthedRequest) => {
  try {
    const data = await backendFetch<MeResponse>("/api/me", { token: req.token });
    return NextResponse.json(data);
  } catch {
    return NextResponse.json({ error: "Failed to fetch user" }, { status: 500 });
  }
});