import {
  AutocompleteProps,
  ChipTypeMap,
  Autocomplete as MuiAutocomplete,
} from "@mui/material";
import {
  Controller,
  ControllerProps,
  FieldPath,
  FieldValues,
} from "react-hook-form";

type ControlledAutoComplete<
  T,
  Multiple extends boolean | undefined,
  DisableClearable extends boolean | undefined,
  FreeSolo extends boolean | undefined,
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
  ChipComponent extends React.ElementType = ChipTypeMap["defaultComponent"]
> = Omit<ControllerProps<TFieldValues, TName>, "render"> &
  AutocompleteProps<T, Multiple, DisableClearable, FreeSolo, ChipComponent>;

export default function Autocomplete<
  T,
  Multiple extends boolean | undefined,
  DisableClearable extends boolean | undefined,
  FreeSolo extends boolean | undefined,
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
  ChipComponent extends React.ElementType = ChipTypeMap["defaultComponent"]
>(
  props: ControlledAutoComplete<
    T,
    Multiple,
    DisableClearable,
    FreeSolo,
    TFieldValues,
    TName,
    ChipComponent
  >
) {
  const { control, name, rules, shouldUnregister, ...autocompleteProps } =
    props;
  const controllerProps = {
    control,
    name,
    rules,
    shouldUnregister,
  };
  return (
    <Controller
      {...controllerProps}
      render={({ field }) => (
        <MuiAutocomplete
          {...autocompleteProps}
          {...field}
          onChange={(_, value) => field.onChange(value)}
        />
      )}
    />
  );
}
