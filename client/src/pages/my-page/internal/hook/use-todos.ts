import { useEffect, useState } from "react";
import { serverFetch } from "../../../../utils/fetch";

type Todo = {
    id: number;
    title: string;
    completed: boolean;
};

export const  useTodos = (todoListID: number) => {
    const [todos,setTodos] = useState<Todo[]>([]);

    const fetchTodos = async () => {
        const res = await serverFetch(`/api/todos?todo_list_id=${todoListID}`);
        setTodos(await res.json());
    };

    useEffect(() => {
        fetchTodos();
    },[]);

    return {
        todos,
        fetchTodos,
    };
};