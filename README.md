# Short URL — URL Shortener with Analytics

> Shorten URLs. Track clicks. Own your data.

---

## Features

| | |
|---|---|
| 🔗 **URL Shortening** | Paste a long URL, get a clean short link instantly |
| 🔐 **Auth** | Register and login to manage your personal links |
| ✅ **Validation** | URL format checking and helpful error feedback |
| 🎯 **Dashboard** | View all your shortened URLs in one place |
| 📊 **Analytics (WIP)** | Track clicks and usage per link from your dashboard |

---

## Pages

- `/` — Landing page
- `/register` — Create an account
- `/login` — Sign in
- `/dashboard` — Manage all your links + create new ones
- `/[slug]` — Redirects to the long url
- `/*` — 404 error page

---

## Architecture

```
snip/
├── backend/          # Go — REST API, auth, URL logic
├── frontend/         # Next.js — UI, routing, forms
└── infra/            # Coming soon — Terraform, CI/CD
```

### Stack

- **Backend** — Go
- **Frontend** — Next.js
- **Analytics** — *(coming soon)*
- **Infra** — *(coming soon — Terraform + GitLab CI/CD)*

---

## Getting Started

### Prerequisites

- Go 1.22+
- Node.js 20+

### Run locally

```bash
# Backend
cd backend
go run main.go

# Frontend
cd frontend
npm install
npm run dev
```

---

## Roadmap

- [x] URL shortening
- [x] User auth
- [x] Dashboard
- [ ] Click analytics
- [ ] Infrastructure as code (Terraform)
- [ ] CI/CD pipeline (GitLab)

---

Built with the help of AI tools.