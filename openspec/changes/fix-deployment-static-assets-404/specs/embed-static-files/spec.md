## MODIFIED Requirements

### Requirement: Static files are served from embedded FS
The system SHALL serve static files from the embedded FS with correct file paths and MIME types.

#### Scenario: CSS file is served with correct MIME type
- **WHEN** a request is made to `/assets/index.b4b296db.css`
- **THEN** the system returns the CSS file with `Content-Type: text/css; charset=utf-8`

#### Scenario: JavaScript file is served with correct MIME type
- **WHEN** a request is made to `/assets/index.abc123.js`
- **THEN** the system returns the JavaScript file with `Content-Type: application/javascript; charset=utf-8`

#### Scenario: Static file is served from correct path
- **WHEN** a request is made to any static file path (e.g., `/assets/...`, `/favicon.ico`)
- **THEN** the system returns the file from the embedded `static/` directory

#### Scenario: manifest.webmanifest is served correctly
- **WHEN** a request is made to `/manifest.webmanifest`
- **THEN** the system returns the manifest file with `Content-Type: application/manifest+json`

### Requirement: Support SPA routing
The system SHALL support Vue Router history mode by serving `index.html` for unmatched routes.

#### Scenario: Unmatched route returns index.html
- **WHEN** a request is made to a non-API route (e.g., `/dashboard`)
- **THEN** the system returns `index.html` to allow Vue Router to handle the route

#### Scenario: API routes are not affected
- **WHEN** a request is made to `/admin/*` or `/api/*`
- **THEN** the system routes to the appropriate API handler
