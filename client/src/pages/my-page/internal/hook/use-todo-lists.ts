import { useEffect, useState } from "react";
import { serverFetch } from "../../../../utils/fetch";

type TodoList = {
    id: number;
    name: string;
}
//useTodoを定義した
export const useTodoLists = () => {
    const [todoLists,setTodoLists] = useState<TodoList[]>([]);

    const fetchTodoLists = async () => {
        const res = await serverFetch("/api/todolists");
        const data = await res.json();
        setTodoLists(data);
    }

    useEffect(() => {
        fetchTodoLists();
    },[])

    return {
        todoLists,
        fetchTodoLists,
    }
}