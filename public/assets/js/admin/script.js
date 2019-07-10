$(document).ready(function(){
    $('.summernote').summernote({
        height: 355,   //set editable area's height
    });

    $(".select2_demo").select2({
        placeholder: "Select a category",
        allowClear: true
    });

    $('.dataTables-example').DataTable({
        pageLength: 25,
        responsive: true,
        dom: '<"html5buttons"B>lTfgitp',
        buttons: [
            { extend: 'copy'},
            {extend: 'csv'},
            {extend: 'excel', title: 'ExampleFile'},
            {extend: 'pdf', title: 'ExampleFile'},

            {extend: 'print',
            customize: function (win){
                    $(win.document.body).addClass('white-bg');
                    $(win.document.body).css('font-size', '10px');

                    $(win.document.body).find('table')
                            .addClass('compact')
                            .css('font-size', 'inherit');
            }
            }
        ]

    });

    $('.js-clear-form').on('click', function() {
        $(this).closest('form').find("input[type=text], textarea").val("");
    });
});
