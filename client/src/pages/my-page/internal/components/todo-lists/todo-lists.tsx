import { Dispatch, SetStateAction } from "react";

import "./todo-lists.css";

type TodoList = {
    id: number;
    name: string;
};

interface TodoListsProps {
    todoLists: TodoList[];
    setTodoListID: Dispatch<SetStateAction<number>>;
}

export const TodoLists = ({ todoLists, setTodoListID }: TodoListsProps) => {

    return(
        <div className="todo-lists__container">
            {todoLists.map((list) => (
                <button
                    key={list.id}
                    onClick={() => setTodoListID(list.id)}
                >{list.name}</button>
            ))}
        </div>
    )
}