enum ApplicationState {
  DRAFT = "DRAFT",
  APPLIED = "APPLIED",
  OFFERED = "OFFERED",
  ACCEPTED = "ACCEPTED",
  REJECTED = "REJECTED",
  WITHDRAWN = "WITHDRAWN",
  CLOSED = "CLOSED",
}

type Application = {
  id: string;
  title: string;
  description: string;
  url: string;
  state: ApplicationState;
  applied_at: Date | null;
  created_at: Date | null;
  updated_at: Date | null;
};

export { Application, ApplicationState };
