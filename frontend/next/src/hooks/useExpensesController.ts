import { useRef, useState } from "react";

export const useExpensesController = ({
  category,
  date,
}: {
  category: string;
  date: Date;
}) => {
  const expenses = useRef<Array<Expense>>([]);
  const [loading, setIsLoading] = useState(true);

  const addExpense = (expense: Expense) => {};
};
