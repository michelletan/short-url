import { NextRequest, NextResponse } from "next/server";

export type AuthedRequest = NextRequest & { token: string };

type RouteHandler = (req: AuthedRequest) => Promise<NextResponse>;

/**
 * Wraps a route handler, ensuring a valid session token exists.
 * Injects the token into the request object so handlers don't need to re-read it.
 *
 * Usage:
 *   export const GET = withAuth(async (req) => {
 *     const data = await backendFetch("/api/links", { token: req.token });
 *     return NextResponse.json(data);
 *   });
 */
export function withAuth(handler: RouteHandler) {
  return async (req: NextRequest): Promise<NextResponse> => {
    const token = req.cookies.get("session")?.value;

    if (!token) {
      return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
    }

    // Inject token so handlers can use it without re-reading the cookie
    (req as AuthedRequest).token = token;

    return handler(req as AuthedRequest);
  };
}