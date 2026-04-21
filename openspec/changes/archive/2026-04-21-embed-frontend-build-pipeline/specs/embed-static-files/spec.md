## ADDED Requirements

### Requirement: Embed static files into binary
The system SHALL embed frontend static files into the Go binary using the `embed` package, enabling single-binary deployment.

#### Scenario: Static files are embedded at build time
- **WHEN** the Go binary is built
- **THEN** all files in `backend/static` directory are embedded into the binary

#### Scenario: Static files are served from embedded FS
- **WHEN** a request is made to the root path `/`
- **THEN** the system serves the embedded `index.html` file

### Requirement: Support SPA routing
The system SHALL support Vue Router history mode by serving `index.html` for unmatched routes.

#### Scenario: Unmatched route returns index.html
- **WHEN** a request is made to a non-API route (e.g., `/dashboard`)
- **THEN** the system returns `index.html` to allow Vue Router to handle the route

#### Scenario: API routes are not affected
- **WHEN** a request is made to `/admin/*` or `/api/*`
- **THEN** the system routes to the appropriate API handler

### Requirement: Exclude static files from version control
The system SHALL exclude the `backend/static` directory from version control via `.gitignore`.

#### Scenario: Static directory is ignored
- **WHEN** files are added to `backend/static`
- **THEN** git does not track these files
