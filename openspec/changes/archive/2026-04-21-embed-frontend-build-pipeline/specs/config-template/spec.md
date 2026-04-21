## ADDED Requirements

### Requirement: Configuration file template
The system SHALL provide a configuration file template with `.example` suffix for users to copy and customize.

#### Scenario: Template file exists
- **WHEN** user clones the repository
- **THEN** a `config.conf.example` file exists in the `backend` directory

#### Scenario: Template contains all required fields
- **WHEN** user opens `config.conf.example`
- **THEN** all configuration fields are documented with example values

### Requirement: Exclude actual config from version control
The system SHALL exclude `backend/config.conf` from version control via `.gitignore`.

#### Scenario: Config file is ignored
- **WHEN** user creates or modifies `backend/config.conf`
- **THEN** git does not track the file

### Requirement: Document configuration setup
The system SHALL document the configuration setup process in README or similar documentation.

#### Scenario: User knows how to configure
- **WHEN** user reads the documentation
- **THEN** instructions explain copying `config.conf.example` to `config.conf` and editing values
