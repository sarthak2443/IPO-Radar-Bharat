# IPO Radar

IPO Radar is a typed React/Vite frontend for discovering Indian IPOs, understanding
subscription demand, and tracking GMP and listing performance.

## Local development

```bash
npm install
npm run dev
```

The production build is generated with `npm run build`.

## GitHub Pages

The repository is configured for the project-site URL:

`https://sarthak2443.github.io/IPO-Radar/`

The Vite base path is enabled only in GitHub Actions, and
`.github/workflows/deploy-pages.yml` builds and deploys `dist` on pushes to `main`.
In the repository settings, set **Pages → Source** to **GitHub Actions**.

## Live data plan

GitHub Pages is static hosting and cannot safely hold provider credentials or run
scheduled ingestion. The frontend therefore uses `src/services/ipoService.ts` as a
provider boundary. A future backend should:

1. Run a scheduled ingestion job (for example, a small Node/TypeScript worker on
   a managed service) that normalizes official exchange data and licensed provider
   feeds into one database.
2. Keep provider adapters separate for IPO master data, subscription, GMP,
   allotment, financials, and listing performance.
3. Expose read-only endpoints such as `GET /api/ipos`, `GET /api/ipos/:id`,
   `GET /api/ipos/:id/subscription`, and `GET /api/ipos/:id/gmp-history`.
4. Return `source`, `updatedAt`, and `availability` with every important metric.
   GMP must be labeled unofficial; missing or delayed values should be explicit.
5. Add caching/rate limits and never scrape InvestorGain directly from the
   browser. InvestorGain is useful as a product benchmark, but its public pages
   do not provide a dependable free API; confirm terms and licensing before using
   any provider feed.

Recommended first backend: a small TypeScript API plus PostgreSQL (or Supabase
Postgres) with a scheduled worker. Start with official NSE/BSE data for dates,
issue details, subscription and listing data; add a licensed GMP provider only
after the product has a clear source agreement.
