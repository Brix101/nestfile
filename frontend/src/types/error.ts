interface ServerError {
  message: string;
  errors?: ServerErrorDetail[];
}

interface ServerErrorDetail {
  field: string;
  message: string;
}
