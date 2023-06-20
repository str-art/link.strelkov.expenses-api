"use client";

import { useForm } from "react-hook-form";
import Autocomplete from "./Autocomplete";
import dayjs from "dayjs";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import addExpense from "@/utils/requests/addExpense";
import useClient from "@/hooks/useClient";
import { useRouter } from "next/navigation";
import useExpenseController from "@/hooks/useExpenseController";
import Loading from "@/app/loading";

export default function ExpenseForm({
  categories,
}: {
  categories: Array<string>;
}) {
  const { control, register, handleSubmit, reset } = useForm<ExpenseInput>({
    defaultValues: {
      category: "Other",
      date: dayjs().format("YYYY-MM-DD"),
      amount: 0,
      name: "",
    },
  });

  const { isLoading, addExpense } = useExpenseController();
  const router = useRouter();

  const onSubmit = async (data: ExpenseInput) => {
    await addExpense(data);
    router.refresh();
    reset();
  };

  return (
    <>
      {isLoading ? (
        <Loading />
      ) : (
        <Card
          sx={{
            width: "100%",
            padding: "1em",
          }}
        >
          <Grid
            container
            component={"form"}
            spacing={2}
            onSubmit={handleSubmit(onSubmit)}
            direction="column"
            width="100%"
            alignItems="center"
          >
            <Grid item minWidth="20em">
              <Autocomplete
                control={control}
                name="category"
                options={categories}
                renderInput={(params) => (
                  <TextField {...params} label="Category" required />
                )}
                freeSolo
                fullWidth
                autoSelect
              />
            </Grid>
            <Grid item>
              <TextField {...register("name")} label="Name" required />
            </Grid>
            <Grid item>
              <Grid
                container
                spacing={2}
                justifyContent="center"
                width="100%"
                wrap="wrap"
              >
                <Grid item>
                  <TextField
                    {...register("amount", { valueAsNumber: true })}
                    type="number"
                    label="Amount"
                    required
                    sx={{
                      maxWidth: "5em",
                    }}
                  />
                </Grid>
                <Grid item>
                  <TextField
                    type="date"
                    {...register("date")}
                    label="Date"
                    required
                  />
                </Grid>
              </Grid>
            </Grid>
            <Grid item>
              <Button variant="contained" type="submit">
                {"Submit"}
              </Button>
            </Grid>
          </Grid>
        </Card>
      )}
    </>
  );
}
