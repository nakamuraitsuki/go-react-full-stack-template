import { useActionState, useCallback } from "react";
import { serverFetch } from "../../../../../utils/fetch";
import { Input } from "../../../../../components/input";
import { Button } from "../../../../../components/button";
import "./todo-create-from.css";

type AddTodoFormStateType = {
    message: string;
};

interface TodoCreateFormProps {
    todoListID: number;
    refetch: () => void;
}

export const TodoCreateForm = ({ todoListID, refetch }: TodoCreateFormProps) => {
    const AddTodoAction = useCallback(
        async (
            _prevState: AddTodoFormStateType,
            formData: FormData
        ): Promise<AddTodoFormStateType> => {
            const title = formData.get("title");

            const res = await serverFetch("/api/todos", {
                method: "POST",
                body: JSON.stringify({
                    todo_list_id: todoListID,
                    title: title,
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
                message: "Todoの追加に失敗しました",
            };
        },
        [refetch]
    );

    const [error, submitAction] = useActionState(AddTodoAction, {
        message: "",
    });
    
    return (
    <div>
        <form action={submitAction} className="TodoCreateForm">
        <Input type="text" name="title" className="TodoCreateForm__input" />
        <Button type="submit">追加</Button>
        </form>
        {error && <p className="TodoCreateForm__error">{error.message}</p>}
    </div>
    );
}