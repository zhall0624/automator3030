import { Typography } from "@mui/material";
import { Step } from "../../types"
import { useQuery } from "@tanstack/react-query";

export default function StepManagement({ id }: { id: number }) {
  const fetchWorkflow = (): Promise<Step> =>
    fetch(`http://localhost:8080/steps/${id}`).then((response) => response.json());

  const { data, isError, isLoading } = useQuery<Step>({
    queryKey: ["steps", id],
    queryFn: fetchWorkflow,
  });

  if (isError) {
    return <Typography>Error Loading Step</Typography>
  }

  if (isLoading) {
    return <Typography>Loading Step</Typography>
  }

  return (
    <>
      <Typography>{data?.name}</Typography>
    </>
  )
}
