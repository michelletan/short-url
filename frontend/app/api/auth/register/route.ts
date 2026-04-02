import { NextRequest, NextResponse } from "next/server";
import { backendFetch } from "@/lib/api";

export async function POST(req: NextRequest) {
  try {
    const { username, email, password } = await req.json();

    if (!username || !email || !password) {
      return NextResponse.json(
        { error: "Username, email, and password are required" },
        { status: 400 }
      );
    }

    const data = await backendFetch<{ token?: string; message?: string }>(
      "/auth/register",
      { method: "POST", body: { username, email, password } }
    );

    return NextResponse.json(data, { status: 201 });
  } catch (err) {
    const message = err instanceof Error ? err.message : "Registration failed";
    return NextResponse.json({ error: message }, { status: 400 });
  }
}