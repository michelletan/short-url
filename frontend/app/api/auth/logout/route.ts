import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  try {
    const cookie = req.headers.get("cookie") ?? "";

    await fetch(
      `${process.env.REDIRECT_BASE_URL ?? "http://localhost:8080"}/auth/logout`,
      {
        method: "POST",
        headers: { Cookie: cookie },
      }
    );

    const response = NextResponse.json({ message: "Logged out" });

    // Clear the session cookie in the browser
    response.cookies.set("session", "", {
      httpOnly: true,
      expires: new Date(0),
      path: "/",
    });

    return response;
  } catch {
    return NextResponse.json({ error: "Logout failed" }, { status: 500 });
  }
}