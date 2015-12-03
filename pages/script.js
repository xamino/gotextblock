function prepare() {
    // define what to do for all dragable elements
    $("[draggable=true]").each(function(index, element){
        // on drag start
        $(element).on("dragstart", function(){
            $("#output").text("Drag started with <"+ $(element).text()+">")
        })
        // on drag end
        $(element).on("dragend", function(){
            $("#output").text("Drag ended with <"+$(element).text()+">")
            // append data to text field?
        })
    })
}

function test() {
    $("#output").text("Hello JS and JQuery World!")
}
