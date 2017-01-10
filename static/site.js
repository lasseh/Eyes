$(document).ready(function() {

	var fetchInterval = 10000;

	status();
	setInterval(status, fetchInterval);

	function status() {
		$.getJSON("/status", function( data ) {
			// Clear alarms for each loop
			$('#status').text('');

			// No active alarms
			if (data.alert.length === 0) {
				var msg = '<div class="panel panel-success"> \
					   <div class="panel-heading"> \
					    <h3 class="panel-title">No Outages</h3> \
					   </div> \
					  <div class="panel-body"> \
					   <center><i class="fa fa-check fa-4x success" aria-hidden="true"></i></center> \
					   </div> \
					  </div>';
				$("#status").html(msg);
			} else {
				// Display active alarms
				$.each(data.alert, function (key, value) {
					var msg = '<div class="panel panel-' + value.warninglevel + '"> \
						  <div class="panel-heading"> \
						   <h2 class="panel-title">' + value.ruleName + ' for ' + value.monitors[0].prefix +'</h2> \
						  </div> \
						  <div class="panel-body"> \
						   <div class="row"> \
						    <div class="col-md-8"> \
						     <p class="h4"> \
						     <i class="fa fa-lg ' + value.warningicon + '" aria-hidden="true"></i> ' + value.message +'</p> \
						     <h5><small>' + value.monitors[0].metricsAtStart +'</small></h5> \
						    </div> \
						   <div class="col-md-4"> \
						    <p class="text-right text-muted h5">Started: ' + value.dateStart+ '</p> \
						   </div> \
						  </div> \
						</div> \
					</div>';
					$("#status").append(msg);
				});
			}
		});
	};
});
