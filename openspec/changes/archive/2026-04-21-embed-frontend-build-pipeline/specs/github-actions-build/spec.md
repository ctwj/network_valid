## ADDED Requirements

### Requirement: Manual build trigger with version input
The GitHub Action workflow SHALL support manual triggering with a version number input.

#### Scenario: User triggers build with version
- **WHEN** user triggers the workflow with version "v1.0.0"
- **THEN** the build artifacts are tagged with version "v1.0.0"

#### Scenario: Default version when not specified
- **WHEN** user triggers the workflow without specifying a version
- **THEN** the build uses a default version format (e.g., "v0.0.1-dev")

### Requirement: Cross-platform build support
The workflow SHALL build binaries for multiple platforms (Linux amd64, Windows amd64).

#### Scenario: Build for Linux
- **WHEN** the workflow runs
- **THEN** a Linux amd64 binary is produced

#### Scenario: Build for Windows
- **WHEN** the workflow runs
- **THEN** a Windows amd64 binary is produced

### Requirement: Frontend build integration
The workflow SHALL build the frontend and copy artifacts to the backend static directory before building the binary.

#### Scenario: Frontend is built before backend
- **WHEN** the workflow runs
- **THEN** frontend is built with `pnpm build`
- **AND** the output is copied to `backend/static`

### Requirement: GitHub Release creation
The workflow SHALL create a GitHub Release with the built binaries as assets.

#### Scenario: Release is created with artifacts
- **WHEN** the build completes successfully
- **THEN** a GitHub Release is created with the specified version
- **AND** the binaries are uploaded as release assets
