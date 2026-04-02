import { NextRequest, NextResponse } from "next/server";
import { backendFetch } from "@/lib/api";

export async function POST(req: NextRequest) {
  try {
    const { email, password } = await req.json();

    if (!email || !password) {
      return NextResponse.json(
        { error: "Email and password are required" },
        { status: 400 }
      );
    }

    const res = await fetch(
      `${process.env.REDIRECT_BASE_URL ?? "http://localhost:8080"}/auth/login`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      }
    );

    if (!res.ok) {
      const error = await res.text();
      return NextResponse.json(
        { error: error || "Invalid credentials" },
        { status: 401 }
      );
    }

    const data = await res.json() as { access_token: string; expires_in: number };
    const response = NextResponse.json({ success: true });
    
    if (data.access_token) {
      response.cookies.set("session", data.access_token, {
        httpOnly: true,
        secure: process.env.NODE_ENV === "production",
        sameSite: "lax",
        path: "/",
        maxAge: data.expires_in, // e.g. 3600 for 1 hour
      });
    }
    
    return response;
  } catch (err) {
    const message = err instanceof Error ? err.message : "Login failed";
    return NextResponse.json({ error: message }, { status: 500 });
  }
}