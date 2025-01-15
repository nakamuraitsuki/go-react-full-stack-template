import { useActionState, useCallback } from "react"
import { serverFetch } from "../../../../../utils/fetch";
import { Input } from "../../../../../components/input";
import { Button } from "../../../../../components/button";

type AddTodoListFormStateType = {
    message: string;
}

interface TodoListsCreateForm {
    refetch: () => void;
}

export const TodoListsCreateForm = ({ refetch }: TodoListsCreateForm) => {
    const AddTodoListAction = useCallback(
        async (
            _prevState: AddTodoListFormStateType,
            formdata: FormData
        ): Promise<AddTodoListFormStateType> => {
            const name = formdata.get("name");

            const res = await serverFetch("/api/todo-lists", {
                method: "POST",
                body: JSON.stringify({
                    name: name,
                }),
                headers: {
                    "Content-Type": "application/json",
                },
            });

            if (res.ok) {
                refetch();
                return { message: "" };
            }

            return {
                message: "Todoの追加に失敗しました"
            };
        },
        [refetch]
    );

    const [error, submitAction] = useActionState(AddTodoListAction, {
        message: "",
    });

    return (
        <div>
            <form action={submitAction}>
            <Input type="text" name="name" />
            <Button type="submit">追加</Button>
            </form>
            {error && <p>{error.message}</p>}
        </div>
    );
}