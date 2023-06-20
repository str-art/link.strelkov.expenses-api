import { useState } from "react";
import useClient from "./useClient";
import { startTransition } from "react";

export default function useExpenseController() {
  const client = useClient();
  const [isLoading, setIsLoading] = useState(false);

  const addExpense = async (expense: ExpenseInput) => {
    startTransition(() => {
      setIsLoading(true);
    });
    await client.post(`/expenses/${expense.category}/${expense.date}`, {
      amount: expense.amount,
      name: expense.name,
    });
    startTransition(() => {
      setIsLoading(false);
    });
  };

  return {
    addExpense,
    isLoading,
  };
}
