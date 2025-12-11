import { CreateApplicationForm } from "@/components/form/create-application";
import { Application } from "@/types/application";

export default async function ApplicationDetailsPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const data = await fetch(`${process.env.API_BASE_URL}/applications/${id}`);
  const application = (await data.json()) as Application;

  return (
    <div>
      <h1>Application Details</h1>

      <CreateApplicationForm />
    </div>
  );
}
